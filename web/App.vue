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
        <template #fallback>
            <div class="fixed left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 z-1">
                <Spinner class="animate-spin w-52 h-52" />
            </div>
        </template>
    </Suspense>
</template>