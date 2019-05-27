<template>
    <div>       
         <paginate name="keys" :list="keys"
            :per="perPage"></paginate>   
            <b-card-group columns>
                 <key-item
                    v-for="key in paginated('keys')"
                    :key="key.IdLlave"
                    :keyItem="key"
                >
                </key-item>                    
            </b-card-group>
            <paginate-links for="keys" :classes="{
                'ul':'pagination',
                'li' : 'paginatepage-item',
                'li > a':'page-link'
            }"
            ></paginate-links>  
    </div>
</template>
<script>
import KeyItem from './KeyItem'
export default {
    components:{
        KeyItem
    },
    props:{
        keys:{
            type:Array,
            required:true
        },
    },
    data(){
        return{
            paginate:['keys'],
            perPage:6
        }
    },
}
</script>

