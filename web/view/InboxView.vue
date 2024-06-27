<script setup>
import useGetQuestions from '@/composable/useGetQuestions';
import Nav from '@/component/Nav.vue';
import QuestionModal from '@/component/QuestionModal.vue'
import LogoutButton from '@/component/LogoutButton.vue'
import InboxNew from '@/icon/InboxNew.vue';
import Inbox from '@/icon/Inbox.vue';
import useSeeQuestion from '@/composable/useSeeQuestion';

const {questions, fetchQuestions, fetchStatus, loading} = useGetQuestions()
const { question, showModal, seeQuestion } = useSeeQuestion()

fetchQuestions()

</script>

<template>
    <div class="w-[100vw] min-h-[100vh] place-items-center fixed bg-gradient-to-br from-blue-400 to-blue-600">
    <Nav />
    <LogoutButton />
    <div class="max-w-2xl m-auto">
        <div v-if="questions.length == 0 && loading == false" class="text-white">data kosong!</div>
        <div v-if="loading">Mengambil data...</div>
        <div v-if="fetchStatus === 'success'" class="grid grid-cols-3 gap-5 place-items-center p-10">
            <button v-for="item in questions" @click="seeQuestion(item)" class="bg-blue-300 rounded-3xl">
                <Inbox v-if="item.seen" class="h-24 w-24" />
                <InboxNew v-else class="h-24 w-24" />
            </button>
        </div>
    </div>
    <QuestionModal title="Ask Me Anything!" :question :show="showModal" @show="(e) => showModal=e" />
    </div>
</template>