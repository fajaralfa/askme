<script setup>
import useGetQuestions from '@/composable/useGetQuestions';
import Nav from '@/component/Nav.vue';
import QuestionModal from '@/component/QuestionModal.vue'
import LogoutButton from '@/component/LogoutButton.vue'
import { ref } from 'vue';

const {questions, fetchQuestions, fetchStatus} = useGetQuestions()

fetchQuestions()

const question = ref('')
const show = ref(false)
function showModal(q) {
    question.value = q
    show.value = true
}

</script>

<template>
    <Nav />
    <LogoutButton />
    <div class="max-w-2xl m-auto">
        <div class="bg-red-500" v-if="questions.length == 0">data kosong</div>
        <div v-if="fetchStatus === 'success'" class="grid grid-cols-3 gap-5 place-items-center p-10">
            <button v-for="item in questions" @click="showModal(item.question)"
                class="bg-blue-500 h-24 w-24 rounded-3xl"></button>
        </div>
    </div>
    <QuestionModal title="Ask Me Anything!" :question :show @show="(e) => show=e" />
</template>