import Vue from 'vue'
export async function fetchKeys ({commit}) {
  try {
    const {data} = await Vue.axios({
      url: '/getRSAKeys'
    })    
    commit('setEncryptedText', '')
    commit('setKeys', data)
    setTimeout(function(){  commit('cleanErrors') }, 3000);
  } catch (e) {      
    commit('keysError', e.message)
  } 
}

export async function addKey ({commit}, key) {  
  try {
    await Vue.axios({
      method: 'POST',
      url: '/createRSAKey',
      data: {        
        NombreLlave: key.NombreLlave,        
      }
    })
    commit('keyCreated',true)
    setTimeout(function(){  commit('keyCreated',false) }, 5000);
  } catch (e) {
    commit('keysError', e.message)
  } 
}

export async function encryptText ({commit}, key) {    
  try {
    const {data} = await Vue.axios({
      method: 'POST',
      url: '/encrypt',
      data: {                     
        IdLlave:key.IdLlave,
        TextoPlano:key.Texto
      }
    })
    commit('setEncryptedText', data)
    commit('cleanErrors','')
  } catch (e) {
    commit('setEncryptedText', e.message)
    commit('keysError', e.message)
  } 
}

export async function decryptText ({commit}, key) {    
  try {
    const {data} = await Vue.axios({
      method: 'POST',
      url: '/decrypt',
      data: {                     
        IdLlave:key.IdLlave,
        TextoPlano:key.TextoEncriptado
      }
    })
    commit('setEncryptedText', data)
    commit('cleanErrors','')
  } catch (e) {
    commit('keysError', e.message)
    commit('setEncryptedText', e.message)
  } 
}

