<script setup>
import useGetQuestions from '@/composable/useGetQuestions';
import Nav from '@/component/Nav.vue';
import QuestionModal from '@/component/QuestionModal.vue'
import LogoutButton from '@/component/LogoutButton.vue'
import InboxNew from '@/icon/InboxNew.vue';
import Inbox from '@/icon/Inbox.vue';
import { ref } from 'vue';
import { fetchJSONWithAuth } from '@/helper/fetch';
import * as user from '@/store/userStore'

const {questions, fetchQuestions, fetchStatus} = useGetQuestions()

await fetchQuestions()

const question = ref('')
const see = ref(false)
async function seeQuestion(item) {
    if (!item.seen) {
        await fetchJSONWithAuth(`/api/v1/questions/${item.id}`, user.get().value.accessToken) // trigger seen status
    }
    question.value = item.question
    item.seen = true
    see.value = true
}

</script>

<template>
    <Nav />
    <LogoutButton />
    <div class="max-w-2xl m-auto">
        <div class="flex flex-col items-center gap-10 mt-28">
            <div class="" v-if="questions.length == 0">data kosong!</div>
        </div>
        <div v-if="fetchStatus === 'success'" class="grid grid-cols-3 gap-5 place-items-center p-10">
            <button v-for="item in questions" @click="seeQuestion(item)" class="bg-blue-300 rounded-3xl">
                <Inbox v-if="item.seen" class="h-24 w-24" />
                <InboxNew v-else class="h-24 w-24" />
            </button>
        </div>
    </div>
    <QuestionModal title="Ask Me Anything!" :question :see @see="(e) => see=e" />
</template>