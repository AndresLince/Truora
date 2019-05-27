<template>    
    <div>
    <b-card 
      title="Crear llave:"        
      tag="article"      
      class="mb-12 mt-4"
    >
      <b-card-text>
      <b-form @submit.prevent="$emit('processKey', keyItem)">
          <b-form-group
              id="keyItem"
              label="Nombre de la llave"
              label-for="keyItem"
          >
          <b-form-input
                  autocomplete="off"
                  id="keyItem"
                  v-model="keyItem.NombreLlave"
                  :state="!$v.keyItem.NombreLlave.$invalid"
                  placeholder="Introduce la tarea"
                  @input="$v.keyItem.$touch"
              ></b-form-input>
              <b-form-invalid-feedback id="todoInfo" v-if="$v.keyItem.$dirty">
                  Este campo es requerido y debe tener una longitud mínima de 4
              </b-form-invalid-feedback>
          </b-form-group>
          <b-button
              type="submit"
              variant="primary"
              :disabled="$v.keyItem.$invalid">
              {{ keySubmit }}
          </b-button>        
      </b-form>
      </b-card-text>    
    </b-card>
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
        },
        keySubmit: {
          type: String,
          default: 'Crear'
        }
      },
      validations: {
        keyItem: {
          NombreLlave: {
            required, minLength: minLength(4)
          },
        }
      }
    }
</script>