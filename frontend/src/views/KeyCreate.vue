<template>
    <div class="mt-2">
        <div v-if="error"> 
          <b-alert show variant="warning">{{errorMessage}}</b-alert> 
        </div>
        <key-form :keyItem="keyItem" @processKey="addKey"></key-form>
    </div>
</template>
<script>
import {mapActions,mapGetters} from 'vuex'
import KeyForm from '@/components/KeyForm'
    export default {
      components: {
        KeyForm
      },
      data () {
        return {
          keyItem: {
            NombreLlave: null,            
          }
        }
      },
      methods:{
          ...mapActions({
              createKey:'keys/addKey'
          }),
          addKey(keyItem){
              this.createKey(keyItem).then(()=>{
                  this.$router.push('/keys')
              })
          }
      },
       computed: {
        ...mapGetters('keys', ['error','errorMessage'])
      }    
    }
</script>