<script setup>
import { RouterView } from 'vue-router'
import Spinner from './icon/Spinner.vue';
import { isLoading } from './store/loadingPageStore';
import { Suspense } from 'vue';

const loading = isLoading()
</script>

<template>
    <Spinner class="animate-spin w-10 h-10 right-2 top-2 fixed" v-show="loading" />
    <Suspense>
        <RouterView />
    </Suspense>
</template>