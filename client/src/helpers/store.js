import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

export default new Vuex.Store({

  state: {
    token: localStorage.getItem('token') || '',
    user : JSON.parse(localStorage.getItem('user')) || {},
    status: ''
  },

  mutations: {
    auth_success(state, token, user) {
      state.token = token
      state.user = user
    },
    logout(state) {
      state.token = ''
    },
    get_logged_user_success(state, user) {
      state.user = user
    },
    create_new_poll_success(state) {
      state.status = 'Poll created'
    },
    update_poll_success(state) {
      state.status = 'Poll updated'

    },
    add_vote_success(state) {
      state.status = 'Vote added'
    },
    create_user_success(state) {
      state.status = 'User created';
    },
    update_user_success(state, user) {
      state.status = 'User updated';
      state.user = user
    }
  },

  actions: {
    login({commit}, user){
      return new Promise((resolve, reject) => {
        axios({url: 'http://localhost:8081/login', data: user, method: 'POST'})
        .then(resp => {
          const token = resp.data
          const user = resp.user
          localStorage.setItem('token', token)
          axios.defaults.headers.common['Authorization'] = token
          commit('auth_success', token, user)
          resolve(resp)
        })
        .catch(err => {
          localStorage.removeItem('token')
          reject(err)
        })
      })
    },

    logout({commit}){
      return new Promise((resolve) => {
        commit('logout')
        localStorage.removeItem('token')
        localStorage.removeItem('user')
        delete axios.defaults.headers.common['Authorization']
        resolve()
      })
    },

    getLoggedUser({commit}) {
      return new Promise((resolve, reject) => {
        axios({url: 'http://localhost:8081/user', method: 'GET', headers: {
          'Authorization': 'Bearer ' + localStorage.getItem('token')
        }})
        .then(resp => {
          const user = resp.data
          localStorage.setItem('user', JSON.stringify(user))
          commit('get_logged_user_success', user)
          resolve(resp)
        })
        .catch(err => {
          console.log(err);
          localStorage.removeItem('user')
          reject(err)
        })
      })
    },

    createPoll({commit}, poll) {
      return new Promise((resolve, reject) => {
        axios({
          url: 'http://localhost:8081/votes',
          method: 'POST',
          headers: {
            'Authorization': 'Bearer ' + localStorage.getItem('token')
          },
          data: poll
        })
        .then(resp => {
          commit('create_new_poll_success')
          resolve(resp)
        })
        .catch(err => {
          console.log(err);
          reject(err)
        })
      })
    },

    updatePoll({commit}, poll) {
      return new Promise((resolve, reject) => {
        axios({
          url: 'http://localhost:8081/votes/' + poll.id,
          method: 'PUT',
          headers: {
            'Authorization': 'Bearer ' + localStorage.getItem('token')
          },
          data: poll
        })
        .then(resp => {
          commit('update_poll_success')
          resolve(resp)
        })
        .catch(err => {
          console.log(err);
          reject(err)
        })
      })
    },

    deletePoll({commit}, id) {
      return new Promise((resolve, reject) => {
        axios({
          url: 'http://localhost:8081/votes/' + id,
          method: 'DELETE',
          headers: {
            'Authorization': 'Bearer ' + localStorage.getItem('token')
          }
        })
        .then(resp => {
          commit('delete_poll_success')
          resolve(resp)
        })
        .catch(err => {
          console.log(err);
          reject(err)
        })
      })
    },

    getPolls() {
      return new Promise((resolve, reject) => {
        axios({
          url: 'http://localhost:8081/votes',
          method: 'GET',
          headers: {
            'Authorization': 'Bearer ' + localStorage.getItem('token')
          }
        })
        .then(resp => {
          resolve(resp)
        })
        .catch(err => {
          console.log(err);
          reject(err)
        })
      })
    },

    addVote({commit}, id) {
      return new Promise((resolve, reject) => {
        axios({
          url: 'http://localhost:8081/vote/' + id,
          method: 'PUT',
          headers: {
            'Authorization': 'Bearer ' + localStorage.getItem('token')
          }
        })
        .then(resp => {
          commit('add_vote_success')
          resolve(resp)
        })
        .catch(err => {
          console.log(err);
          reject(err)
        })
      })
    },

    createUser({commit}, user) {
      return new Promise((resolve, reject) => {
        axios({
          url: 'http://localhost:8081/users',
          method: 'POST',
          data: user
        })
        .then(resp => {
          commit('create_user_success')
          resolve(resp)
        })
        .catch(err => {
          console.log(err);
          reject(err)
        })
      })
    },

    updateUser({commit}, user) {
      return new Promise((resolve, reject) => {
        axios({
          url: 'http://localhost:8081/users/' + user.id,
          method: 'PUT',
          headers: {
            'Authorization': 'Bearer ' + localStorage.getItem('token')
          },
          data: user
        })
        .then(resp => {
          commit('update_user_success', resp.data)
          localStorage.setItem('user', JSON.stringify(resp.data))
          resolve(resp)
        })
        .catch(err => {
          console.log(err);
          reject(err)
        })
      })
    }
  },

  getters : {
    isLoggedIn: state => !!state.token,
    loggedUser: state => state.user
  }
})
