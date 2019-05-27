<template>
    <div class="mt-3">
        <b-row> 
            <b-col cols="12">        
                <h2>Listado de Llaves</h2>
                    <div v-if="error"> 
                        <b-alert show variant="warning">{{errorMessage}}</b-alert> 
                    </div>
                     <div v-if="keyCreated"> 
                        <b-alert show variant="success">Se creo la llave correctamente</b-alert> 
                    </div>
                    <div v-if="keys.length">  
                        <b-form-group>
                        <b-form-input
                            id="query"
                            v-model="query"
                            autocomplete="off"
                            placeholder="Escribe el nombre de la llave"
                        >
                        </b-form-input>
                        </b-form-group>
                        <key-list :keys="filteredKeys"></key-list>                            
                    </div>
                    <b-alert show variant="info" v-else>No hay llaves para mostrar</b-alert>                 
            </b-col>              
        </b-row>        
    </div>
</template>
<script>
    import {mapActions, mapGetters} from 'vuex'
    import KeyList from '@/components/KeyList'
    export default {
    data(){
        return{
            query:'',            
        }
    },
    components: {
        KeyList
    },
    mounted () {    
        this.fetchKeys()
    },
    methods: {
        ...mapActions('keys', ['fetchKeys'])
    },
    computed: {
        ...mapGetters('keys', ['keys','error','errorMessage','keyCreated']),
        filteredKeys:function(){
            return this.keys.filter((key)=>{
                return key.NombreLlave.includes(this.query)
            })
        }
        
    }  
    }
</script>
