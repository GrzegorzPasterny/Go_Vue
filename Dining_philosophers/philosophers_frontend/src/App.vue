<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Welcome on a philosophers dinner</h1>
        <hr/>

        <form v-on:submit.prevent="StartDinner">
          <div>
            <h2>Philosopher 1</h2>
            <div class="form-group">
              <label>Name</label>
              <input type="text" class="form-control" v-model="posts[0].Name"/>
            </div>
            <div class="form-group">
              <label>Time to eat</label>
              <input type="int" class="form-control" v-model="posts[0].TimeToEat"/>
            </div>
            <div class="form-group">
              <label>Time to think</label>
              <input type="int" class="form-control" v-model="posts[0].TimeToThink"/>
            </div>
            <div class="form-group">
              <label>Number of dishes to be eaten</label>
              <input type="int" class="form-control" v-model="posts[0].HowManyDishesToBeEaten"/>
            </div>
          </div>
          <div>
            <h2>Philosopher 2</h2>
            <div class="form-group">
              <label>Name</label>
              <input type="text" class="form-control" v-model="posts[1].Name"/>
            </div>
            <div class="form-group">
              <label>Time to eat</label>
              <input type="int" class="form-control" v-model="posts[1].TimeToEat"/>
            </div>
            <div class="form-group">
              <label>Time to think</label>
              <input type="int" class="form-control" v-model="posts[1].TimeToThink"/>
            </div>
            <div class="form-group">
              <label>Number of dishes to be eaten</label>
              <input type="int" class="form-control" v-model="posts[1].HowManyDishesToBeEaten"/>
            </div>
          </div>
          <div>
            <h2>Philosopher 3</h2>
            <div class="form-group">
              <label>Name</label>
              <input type="text" class="form-control" v-model="posts[2].Name"/>
            </div>
            <div class="form-group">
              <label>Time to eat</label>
              <input type="int" class="form-control" v-model="posts[2].TimeToEat"/>
            </div>
            <div class="form-group">
              <label>Time to think</label>
              <input type="int" class="form-control" v-model="posts[2].TimeToThink"/>
            </div>
            <div class="form-group">
              <label>Number of dishes to be eaten</label>
              <input type="int" class="form-control" v-model="posts[2].HowManyDishesToBeEaten"/>
            </div>
          </div>
          <br> <br> <hr>
          <div class="form-group">
            <button class="btn btn-primary" type="submit">Start!</button>
          </div>
        </form>
      </div>
    </div>
    <div>
        <h1>
            Philosophers
        </h1>
        <table>
            <tr>
                <td > Name </td>
                <td > Status </td>
                <td> Timestamp </td>
            </tr>
            <tr v-for="item in info" :key="item.id">
                <td> {{item.Name}} </td>
                <td> {{item.Status}} </td>
                <td> {{item.TimeStamp}} </td>
            </tr>
        </table>
    </div>
  </div>
</template>

<script>
//import philosophers from './components/Philosophers'
import axios from 'axios'
export default {
  name: 'App', 
  // components:
  // {
  //   philosophers
  // },
  data() {
    return {
      // post1:
      //   {
      //   name1:null,
      //   timeToEat1:2,
      //   timeToThink1:3,
      //   howManyDishesToBeEaten1:3
      // },
      // post2:
      //   {
      //   name2:"Basia",
      //   timeToEat2:1,
      //   timeToThink2:3,
      //   howManyDishesToBeEaten2:2
      // },
      // post3:
      //   {
      //   name3:"Ela",
      //   timeToEat3:2,
      //   timeToThink3:3,
      //   howManyDishesToBeEaten3:2
      // },
      // post1:
      //   {
      //   name1:null,
      //   timeToEat1:null,
      //   timeToThink1:null,
      //   howManyDishesToBeEaten1:null
      // },
      // post2:
      //   {
      //   name2:null,
      //   timeToEat2:null,
      //   timeToThink2:null,
      //   howManyDishesToBeEaten2:null
      // },
      // post3:
      //   {
      //   name3:null,
      //   timeToEat3:null,
      //   timeToThink3:null,
      //   howManyDishesToBeEaten3:null
      // },
      posts: [{
        Name:null,
        TimeToEat:0,
        TimeToThink:0,
        HowManyDishesToBeEaten:0
      },
      {
        Name:null,
        TimeToEat:0,
        TimeToThink:0,
        HowManyDishesToBeEaten:0
      }, 
      {
        Name:null,
        TimeToEat:0,
        TimeToThink:0,
        HowManyDishesToBeEaten:0
      }],
      info: []
    }
  },
  methods: {
    StartDinner() {

      // const input = [{
      //   Name: "Marek",
      //   TimeToEat: 5,
      //   TimeToThink: 4,
      //   HowManyDishesToBeEaten: 4
      // },
      // {
      //   Name: "Basia",
      //   TimeToEat: 2,
      //   TimeToThink: 4,
      //   HowManyDishesToBeEaten: 3
      // },
      // {
      //     Name: "Piotr",
      //     TimeToEat: 6,
      //     TimeToThink: 4,
      //     HowManyDishesToBeEaten: 8
      // }]
      //this.posts = this.post1 + this.post2 + this.post3
      //console.warn(this.post1)
      console.warn(this.posts)
      axios.post("http://localhost:3000/DiningPhilosophers",
        // { Name: this.name,
        // TimeToEat: this.TimeToEat,
        // TimeToThink: this.TimeToThink,
        // HowManyDishesToBeEaten: this.HowManyDishesToBeEaten
        // }
        this.posts
      )
      .then(response =>{
        console.warn(response)
        this.info = response.data
      })
      .catch((error) => {
        window.alert(`The Go Lang server returned an error: ${error}`)
    });
  }
}
}
</script>

<style>
/* #app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
} */
</style>
