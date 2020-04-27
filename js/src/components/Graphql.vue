<template>
  <div>
    <h1>grapql</h1>
    <input type="text" v-model="name">
    <input type="button" value="Add" @click="post()">
    <ul>
      <li v-for="t in tasks" :key="t.id">
        <a>
          <input type="button" value="🗑" @click="del(t)">
          <input type="checkbox" :checked="t.Done" @click="put(t, $event)">
          {{t.Name}}
        </a>
      </li>
    </ul>
    <ul>
    </ul>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Grapql',
  data() {
    return {
      tasks: [],
      name: "",
    }
  },
  created() {
    this.get()
  },
  methods: {
    get() {
      //axios.get('http://localhost:1323/tasks').then((r) => {
      //  this.tasks = r.data
      //})
      axios({
        url: 'http://localhost:1323/graphql/tasks',
        headers: {
          //Authorization: `bearer ${accessToken}`,
          //Accept: 'application/vnd.github.v4.idl'
        },
        method: 'POST',
        data: {
          query: `query {
            GetQuery {
              tasks {
                id
              }
            }
          }`
        }
      })
        .then(res => res.data)
        .then(console.log)

    },
    post() {
      //axios.post('http://localhost:1323/tasks', {name: this.name}).then(() => {
      //  this.name = ""
      //  this.get()
      //})
      axios({
        url: 'http://localhost:1323/graphql/tasks',
        headers: {
          //Authorization: `bearer ${accessToken}`,
          //Accept: 'application/vnd.github.v4.idl'
        },
        method: 'POST',
        data: {
          query: `postQuery {
            postTask() {
              id
            }
          }`
        }
      })
        .then(res => res.data)
        .then(console.log)
    },
    put(t, e) {
      const done = e.target.checked
      axios.put('http://localhost:1323/tasks', {id: t.Id, done: done}).then(() => {
        this.get()
      })
    },
    del(t) {
      console.log(t)
      axios.delete('http://localhost:1323/tasks', {data:{id: t.Id}}).then(() => {
        this.get()
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
