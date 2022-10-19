<template>

  <button type="button" class="btn btn-outline-primary" @click="showModalRunTodoCheck">Show</button>

  <ViewModalRunTodoCheck ref="modalRunTodoCheck" @submit="reload"></ViewModalRunTodoCheck>

</template>

<script setup>
import swal from "sweetalert2";
import {ref} from "vue";

import ViewModalRunTodoCheck from "./ModalInput.vue";
const modalRunTodoCheck = ref()
const showModalRunTodoCheck = () => modalRunTodoCheck.value.showModal()

const reload = () => swal.fire({ icon: 'success', title: 'Success', text: 'Usecase executed', })

</script>

<style scoped>

</style>