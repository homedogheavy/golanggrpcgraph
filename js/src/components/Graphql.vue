<template>
  <div>
    <h1>grapql</h1>
    <input type="text" v-model="name">
    <input type="button" value="Add" @click="post()">
    <ul>
      <li v-for="t in tasks" :key="t.id">
        <a>
          <input type="button" value="🗑" @click="del(t)">
          <input type="checkbox" :checked="t.done" @click="put(t, $event)">
          {{t.name}}
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
      axios({
        url: 'http://localhost:1323/graphql/tasks',
        headers: {},
        method: 'POST',
        data: {
          query: `query {
            GetQuery {
              tasks {
                id,
                name,
                done
              }
            }
          }`
        }
      })
      .then(res => res.data)
      .then(res => {
        this.tasks = res.data.GetQuery.tasks
      })
    },
    post() {
      axios({
        url: 'http://localhost:1323/graphql/tasks',
        headers: {},
        method: 'POST',
        data: {
          query: `mutation {
            PostTask(Name: "${this.name}") {
              id
            }
          }`
        }
      })
      .then(() => {
        this.get()
      })
    },
    put(t, e) {
      const done = e.target.checked
      axios({
        url: 'http://localhost:1323/graphql/tasks',
        headers: {},
        method: 'POST',
        data: {
          query: `mutation {
            PutTask(Id: ${t.id}, Done: ${done}) {
              id
            }
          }`
        }
      })
      .then(() => {
        this.get()
      })
    },
    del(t) {
      axios({
        url: 'http://localhost:1323/graphql/tasks',
        headers: {},
        method: 'POST',
        data: {
          query: `mutation {
            DeleteTask(Id: ${t.id}) {
              id
            }
          }`
        }
      })
      .then(() => {
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
