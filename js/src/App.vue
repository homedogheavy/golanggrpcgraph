<template>
  <div id="app">
    <Http />
    <Grpc />
    <Graphql />
  </div>
</template>

<script>
import Http from './components/Http.vue'
import Grpc from './components/Grpc.vue'
import Graphql from './components/Graphql.vue'

export default {
  name: 'App',
  components: {
    Http,
    Grpc,
    Graphql
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
