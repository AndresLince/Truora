<template>
  <div>
    <b-row>
      <b-col cols="6">
<b-card 
      title="Encriptar Texto:"        
      tag="article"      
      class="mb-12 mt-4">
      <b-card-text>
      <b-form @submit.prevent="$emit('encryptKey', keyItem)">
        <b-form-group            
            label="Nombre de la llave"
            label-for="keyItem_NombreLlave">
            <b-form-input
                autocomplete="off"
                id="keyItem_NombreLlave"
                v-model="keyItem.NombreLlave"                                
                @input="$v.keyItem.$touch"
                disabled
            ></b-form-input>           
        </b-form-group>
        <b-form-group 
            label-for="keyItem_Texto">
            <b-form-textarea
                rows="1"
                autocomplete="off"
                max-rows="6"    
                id="keyItem_Texto"
                v-model="keyItem.Texto"
                :state="!$v.keyItem.Texto.$invalid"
                placeholder="Escribe el texto a encriptar"
                @input="$v.keyItem.$touch"                
            ></b-form-textarea>
            <b-form-invalid-feedback v-if="$v.keyItem.$dirty">
                Este campo es requerido y debe tener una longitud mínima de 4
            </b-form-invalid-feedback>
        </b-form-group>  
         
        <b-button
            type="submit"
            variant="primary"
            :disabled="$v.keyItem.$invalid">
           encriptar
        </b-button>      
         </b-form>          
      </b-card-text>    
    </b-card>
      </b-col>
      <b-col cols="6">
        <b-card 
      title="Desencriptar Texto:"        
      tag="article"      
      class="mb-12 mt-4">
      <b-card-text>
        <b-form @submit.prevent="$emit('decryptKey', keyItem)">                
         <b-form-group
            id="keyItem"            
            label-for="keyItem">
            <b-form-textarea
                autocomplete="off"
                id="keyItem"
                v-model="keyItem.TextoEncriptado"  
                @input="$v.keyItem.$touch"  
                placeholder="Escribe el texto a desencriptar"
                max-rows="6"              
            ></b-form-textarea>
            <b-form-invalid-feedback id="todoInfo" v-if="$v.keyItem.$dirty">
                Este campo es requerido y debe tener una longitud mínima de 4
            </b-form-invalid-feedback>
        </b-form-group> 
        <b-button
            type="submit"
            variant="primary"
              :disabled="$v.keyItem.$invalid">
           Desencriptar 
        </b-button>        
     
    </b-form>
      </b-card-text>    
    </b-card>
      </b-col>

    </b-row>
    
      

      
</div>    
</template>

<script>
    import { validationMixin } from 'vuelidate'
    import { required, minLength } from 'vuelidate/lib/validators'
    export default {
      mixins: [validationMixin],
      props: {
        keyItem: {
          type: Object,
          required: true
        }     
      },
      validations: {
        keyItem: {
          NombreLlave: {
            required, minLength: minLength(4)
          },
          Texto: {
            required, minLength: minLength(4)
          },         
        }
      }
    }
</script>