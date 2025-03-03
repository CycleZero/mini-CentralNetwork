<script setup>
import { ref } from 'vue';
import axios from 'axios';
const fileData=ref(null);
const filelist=ref([]);



const getFiles=async()=>{
    axios.get('http://localhost:45689/file/list').then(res=>{
        filelist.value=res.data;
        console.log(filelist.value);
    })
}
</script>


<template>
    <div class="file-view">
      <ul>
        <li v-for="file in filelist" >
            {{ file.Name }}
        </li>
      </ul>
    </div>
    <el-button type="primary" @click=getFiles()>文件列表</el-button>
  </template>


<style>

</style>