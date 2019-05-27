<template>
  <div>
      <div v-if="selectedKey">  
          <div v-if="error"> 
            <b-alert show variant="warning">{{errorMessage}}</b-alert> 
          </div>
           <b-row>
            <b-col cols="12">        
            <key-encrypt-form 
              :keyItem="selectedKey" 
              @encryptKey="encryptTextFunction" 
              @decryptKey="decryptTextFunction"             
              mostrar:false            
            ></key-encrypt-form>     
         </b-col>  
        </b-row>
        <b-row>
        <b-col cols="12">
          <b-form-textarea
              autocomplete="off"
              id="keyItem"
              v-model="encryptedText"  
              max-rows="6">
          </b-form-textarea>
        </b-col>  
        </b-row>    
      </div>  
  </div>     
</template>

<script>
  import {mapActions, mapState,mapGetters} from 'vuex'
  import KeyEncryptForm from '@/components/KeyEncryptForm'
  export default {
    components: {
      KeyEncryptForm
    },
    computed: {
      ...mapState('keys', ['selectedKey','encryptedText']),
      ...mapGetters('keys', ['error','errorMessage'])
    },
    methods: {
      ...mapActions({
          encryptText:'keys/encryptText',
          decryptText:'keys/decryptText'
      }),
      encryptTextFunction(selectedKey){
          this.encryptText(selectedKey).then(()=>{
              //this.$router.push('/keys')
          })
      },      
      decryptTextFunction(selectedKey){
          this.decryptText(selectedKey).then(()=>{
              //this.$router.push('/keys')
          })
      },     
    }
  }
</script>