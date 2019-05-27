<template>
    <b-row class="mb-3">        
        <b-card
            :title="keyItem.NombreLlave"
            style="max-width:50rem;"
            class="mb-2">
            <b-button 
                variant="primary"
                @click="goToEncryptKey()">Encriptar o Desencriptar Texto
            </b-button>   
        </b-card>      
    </b-row>
</template>
<script>
import {mapMutations,mapActions} from 'vuex'
export default {
    props:{
        keyItem:{
            type:Object,
            required:true
        }
    },    
    methods:{
        goToEncryptKey(){
            this.setKey(this.keyItem)
            this.$router.push({
                name:'keys-encrypt',
                params:{id:this.keyItem.NombreLlave}
            })
        },
        ...mapMutations('keys',['setKey'])             
    }
}
</script>
