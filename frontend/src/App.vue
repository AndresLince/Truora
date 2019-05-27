<template>
  <div id="app">
   <b-container>
      <b-nav tabs>
        <b-nav-item exact to="/keys">Listado de keys</b-nav-item>
        <b-nav-item exact to="/keys/create">Crear Llaves</b-nav-item>        
      </b-nav>
      <router-view/>      
    </b-container>    
  </div>
</template>