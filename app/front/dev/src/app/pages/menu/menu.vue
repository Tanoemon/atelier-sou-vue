<style lang="sass" src="menu.scss"></style>
<script lang="es6" src="menu.js"></script>

<template>

  <div id="v-menu">
    <div class="menu-group">
      <div>
        <div class="title">Basic Menu</div>
        <div class="content">
          <template v-repeat="menus">
            <div class="menu">
              <div class="name">
                <div>{{name}}</div>
                <div>({{note}})</div>
              </div>
              <div class="line"></div>
              <div class="price">{{price}}</div>
            </div>
          </template>
        </div>
      </div>
    </div>
    <div class="menu-group">
    </div>
    <div class="menu-group">
    </div>
  </div>

</template>
