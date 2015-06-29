package hello

import (
	"fmt"
  "io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/cloud/storage"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/api/dogs", handler)
	http.HandleFunc("/api/dogs/", handler3)
	http.HandleFunc("/api/upload", upload)
}

var bucket = "atelier-sou"

type demo struct {
        c   context.Context
        w   http.ResponseWriter
        ctx context.Context
        // cleanUp is a list of filenames that need cleaning up at the end of the demo.
        cleanUp []string
        // failed indicates that one or more of the demo steps failed.
        failed bool
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	hc := &http.Client{
		Transport: &oauth2.Transport{
	    Source: google.AppEngineTokenSource(c, storage.ScopeFullControl),
	    Base:   &urlfetch.Transport{Context: c},
    },
	}
	ctx := cloud.NewContext(appengine.AppID(c), hc)
	d := &demo{
                c:   c,
                w:   w,
                ctx: ctx,
        }

	d.listBucket()
}

func upload(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	hc := &http.Client{
		Transport: &oauth2.Transport{
	    Source: google.AppEngineTokenSource(c, storage.ScopeFullControl),
	    Base:   &urlfetch.Transport{Context: c},
    },
	}
	ctx := cloud.NewContext(appengine.AppID(c), hc)

	err := r.ParseMultipartForm(1048576);
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	wc := storage.NewWriter(ctx, bucket, "slide/" + fileHeader.Filename)
	wc.ContentType = http.DetectContentType(data);

	if _, err := wc.Write(data); err != nil {
		io.WriteString(w, err.Error())
		return
	}

	if err := wc.Close(); err != nil {
		io.WriteString(w, err.Error())
		return
	}
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "dogsのリストe")
}

func (d *demo) errorf(format string, args ...interface{}) {
        d.failed = true
        log.Errorf(d.c, format, args...)
}

func (d *demo) dumpStats(obj *storage.Object) {
        fmt.Fprintf(d.w, "(filename: /%v/%v, ", obj.Bucket, obj.Name)
        fmt.Fprintf(d.w, "ContentType: %q, ", obj.ContentType)
        fmt.Fprintf(d.w, "ACL: %#v, ", obj.ACL)
        fmt.Fprintf(d.w, "Owner: %v, ", obj.Owner)
        fmt.Fprintf(d.w, "ContentEncoding: %q, ", obj.ContentEncoding)
        fmt.Fprintf(d.w, "Size: %v, ", obj.Size)
        fmt.Fprintf(d.w, "MD5: %q, ", obj.MD5)
        fmt.Fprintf(d.w, "CRC32C: %q, ", obj.CRC32C)
        fmt.Fprintf(d.w, "Metadata: %#v, ", obj.Metadata)
        fmt.Fprintf(d.w, "MediaLink: %q, ", obj.MediaLink)
        fmt.Fprintf(d.w, "StorageClass: %q, ", obj.StorageClass)
        if !obj.Deleted.IsZero() {
                fmt.Fprintf(d.w, "Deleted: %v, ", obj.Deleted)
        }
        fmt.Fprintf(d.w, "Updated: %v)\n", obj.Updated)
}

func (d *demo) listBucket() {
	io.WriteString(d.w, "\nListbeeeeucket result:\n")

	query := &storage.Query{Prefix: "slide/CB"}
	for query != nil {
		objs, err := storage.ListObjects(d.ctx, bucket, query)
    if err != nil {
			d.errorf("listBucket: unable to list bucket %q: %v", bucket, err)
	    return
    }
    query = objs.Next

    for _, obj := range objs.Results {
			io.WriteString(d.w, obj.Name)
    }
	}
}
