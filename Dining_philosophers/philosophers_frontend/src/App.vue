<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Welcome on a philosophers dinner</h1>
        <hr/>

        <form v-on:submit.prevent="StartDinner">
          <div class="form-group">
            <label>Name</label>
            <input type="text" class="form-control" v-model="name"/>
          </div>
          <div class="form-group">
            <label>Time to eat</label>
            <input type="int" class="form-control" v-model="timeToEat"/>
          </div>
          <div class="form-group">
            <label>Time to think</label>
            <input type="int" class="form-control" v-model="timeToThink"/>
          </div>
          <div class="form-group">
            <label>Number of dishes to be eaten</label>
            <input type="int" class="form-control" v-model="howManyDishesToBeEaten"/>
          </div>
          <div class="form-group">
            <button class="btn btn-primary" type="submit">Start!</button>
          </div>
        </form>
      </div>
    </div>
    <div>
        <!-- <h1>
            Philosophers
        </h1>
        <table>
            <tr>
                <td> Name </td>
                <td> Status </td>
                <td> Timestamp </td>
            </tr>
            <tr v-for="item in list" v-bind:key="item.id">
                <td> {{item.Name}} </td>
                <td> {{item.Status}} </td>
                <td> {{item.TimeStamp}} </td>
            </tr>
        </table> -->
        <div id="App">
          {{info}}
        </div>
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
      info: []
    }
    // name: '',
    // timeToEat: 1,
    // timeToThink: 1,
    // howManyDishesToBeEaten: 1,

  },
  methods: {
    StartDinner() {

      const input = [{
        Name: "Marek",
        TimeToEat: 5,
        TimeToThink: 4,
        HowManyDishesToBeEaten: 4
      },
      {
        Name: "Basia",
        TimeToEat: 2,
        TimeToThink: 4,
        HowManyDishesToBeEaten: 3
      },
      {
          Name: "Piotr",
          TimeToEat: 6,
          TimeToThink: 4,
          HowManyDishesToBeEaten: 8
      }]
      console.log(input)
      axios.post("http://localhost:3000/DiningPhilosophers",
        // { Name: this.name,
        // TimeToEat: this.TimeToEat,
        // TimeToThink: this.TimeToThink,
        // HowManyDishesToBeEaten: this.HowManyDishesToBeEaten
        // }
        input
      )
      .then(response =>{
        this.info = response
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
