<template>

  <button type="button" class="btn btn-outline-primary" @click="showModalRunTodoCreate">Show</button>

  <ViewModalRunTodoCreate ref="modalRunTodoCreate" @submit="reload"></ViewModalRunTodoCreate>

</template>

<script setup>
import swal from "sweetalert2";
import {ref} from "vue";

import ViewModalRunTodoCreate from "./ModalInput.vue";
const modalRunTodoCreate = ref()
const showModalRunTodoCreate = () => modalRunTodoCreate.value.showModal()

const reload = () => swal.fire({ icon: 'success', title: 'Success', text: 'Usecase executed', })

</script>

<style scoped>

</style>