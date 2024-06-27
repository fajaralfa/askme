import { ref } from 'vue';
import { fetchJSONWithAuth } from '@/helper/fetch';
import * as user from '@/store/userStore'

export default function useSeeQuestion() {
    const question = ref('')
    const showModal = ref(false)
    async function seeQuestion(item) {
        if (!item.seen) {
            await fetchJSONWithAuth(`/api/v1/questions/${item.id}`, user.get().value.accessToken) // trigger seen status
        }
        question.value = item.question
        item.seen = true
        showModal.value = true
    }

    return { question, showModal, seeQuestion }
}