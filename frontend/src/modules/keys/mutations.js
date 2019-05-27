export function setKeys (state, keys) {
    state.keys = keys;
}

export function keysError (state, payload) {    
    state.error = true
    state.errorMessage = payload
    state.keys = []
}

export function keyCreated(state,payload){
    state.keyCreated=payload
}

export function cleanErrors(state) {    
    state.error = false
    state.errorMessage = ''    
}

export function setKey (state, key) {    
    state.selectedKey = key
}

export function setEncryptedText(state, text) {    
    state.encryptedText = text
}