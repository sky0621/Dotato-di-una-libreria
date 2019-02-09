import * as types from './types'

export const strict = false

export const state = () => ({
  notification: {
    activate: false,
    message: '',
    color: ''
  }
})

export const getters = {
  notification: (state) => {
    return state.notification
  }
}

export const mutations = {
  setNotification(state, data) {
    const cdata = _.clone(data)
    state.notification = {
      activate: cdata.activate,
      message: cdata.message,
      color: cdata.color
    }
  }
}

export const actions = {
  // 通知を活性化する
  activateNotification({ commit }, data) {
    commit(types.SET_NOTIFICATION, { activate: true, message: data.message, color: data.color })
  },

  // INFO通知を活性化する
  activateInfoNotification({ commit }, data) {
    commit(types.SET_NOTIFICATION, { activate: true, message: data.message, color: 'green' })
  },

  // ERROR通知を活性化する
  activateErrorNotification({ commit }, data) {
    commit(types.SET_NOTIFICATION, { activate: true, message: data.message, color: 'red' })
  },

  // 通知を非活性化する
  deactivateNotification({ commit }) {
    commit(types.SET_NOTIFICATION, { activate: false, message: '', color: '' })
  }
}
