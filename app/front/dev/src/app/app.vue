<style lang="sass" src="./styles/app.scss"></style>
<script lang="es6" src="./js/app.js"></script>

<template>

<div id="menu" class="back">
  <v-menu></v-menu>
</div>
<div id="staff" class="back"></div>
<div id="concept" class="back"></div>
<div id="access" class="back"></div>
<div id="logo">
  <span id="atelier">ATELIER</span><span id="sou_j"> 想 </span><span id="sou">SOU</span>
</div>
<div id="site_menu">
  <div id="m_menu" class="menuelem" v-on="click: muClick">メニュー</div>
  <div id="m_staff" class="menuelem" v-on="click: sfClick">スタッフ</div>
  <div id="m_concept" class="menuelem" v-on="click: ctClick">コンセプト</div>
  <div id="m_access" class="menuelem" v-on="click: asClick">アクセス</div>
  <div id="m_resavation" class="menuelem">
    <div id="yoyaku">予約</div>
    <div id="tel"></div>
  </div>
</div>

</template>
