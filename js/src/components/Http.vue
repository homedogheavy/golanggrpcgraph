<template>
  <div>
    <h1>http</h1>
    <input type="text">
    <input type="button" value="Add" @click="post()">
    <ul>
      <li v-for="t in tasks" :key="t"><a href="https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-eslint" target="_blank" rel="noopener">{{t}}</a></li>
    </ul>
    <ul>
    </ul>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Http',
  data() {
    return {
      tasks: ["babel1","eslint2","babel3","eslint4"]
    }
  },
  methods: {
    post() {
      axios.get('http://localhost:8080/').then((r) => {
        console.log(r)
      })
    }
  }
}
</script>

<style scoped>
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
