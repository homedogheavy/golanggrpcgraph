<template>
  <div>
    <h1>grpc</h1>
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
const {TaskServiceClient, GetTaskRequest, PostTaskRequest, PutTaskRequest, DeleteTaskRequest} = require('../../proto/task_service_grpc_web_pb.js');
const enableDevTools = window.__GRPCWEB_DEVTOOLS__ || (() => {});
const client = new TaskServiceClient('http://localhost:1323');
enableDevTools([
  client,
]);


export default {
  name: 'Grpc',
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
      const request = new GetTaskRequest();
      client.getTasks(request, {}, (err, response) => {
        this.tasks = response.toObject().tasksList
      });
    },
    post() {
      const request = new PostTaskRequest();
      request.setName(this.name)
      client.postTask(request, {}, (err, response) => {
        console.log(response)
        console.log(err)
        this.name = ""
        this.get()
      });
    },
    put(t, e) {
      const done = e.target.checked
      const request = new PutTaskRequest();
      request.setId(t.id)
      request.setDone(done)
      client.putTask(request, {}, (err, response) => {
        console.log(response)
        console.log(err)
        this.get()
      });
    },
    del(t) {
      const request = new DeleteTaskRequest();
      request.setId(t.id)
      client.deleteTask(request, {}, (err, response) => {
        console.log(response)
        console.log(err)
        this.get()
      });
    },
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
