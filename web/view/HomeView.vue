<script setup>
import { Teleport, Transition, ref } from 'vue';
import Copy from '@/icon/Copy.vue'
import Nav from '@/component/Nav.vue';
import useUserStore from '@/composable/useUserStore';
import LogoutButton from '@/component/LogoutButton.vue'

const user = useUserStore()

const link = `${window.location.origin}/${user.get().value.email}`

async function copyToClipboard(str) {
    await navigator.clipboard.writeText(str)
    toggle.value = true
}

const toggle = ref(false)

</script>

<template>
    <Nav></Nav>
    <LogoutButton />
    <div class="text-base mt-5 p-3 border rounded inline box-border active:outline-white active:outline active:outline-1">
        <span>{{ link }}</span> <Copy @click="copyToClipboard(link)" width="1.5rem" height="1.5rem" style="display: inline;" /></div>
    <Teleport to="body">
        <Transition
        enter-from-class="opacity-0"
        enter-active-class="transition-opacity duration-150 ease"
        leave-active-class="transition-opacity duration-150 ease"
        leave-from-class="opacity-0"
        leave-to-class="opacity-0"
        class=""
        >
            <div v-if="toggle" @click="toggle = false"
                class="fixed top-0 bg-[rgba(0,0,0,0.3)] w-screen h-screen flex justify-center items-center transition">
                <div class="bg-white text-black px-4 py-2 rounded">link disalin!</div>
            </div>
        </Transition>
    </Teleport>
</template>