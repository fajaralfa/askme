import { ref } from "vue"
import { fetchJSONWithAuth } from "@/helper/fetch";
import * as user from '@/store/userStore'

export default function useGetQuestions() {
    const questions = ref([])
    const loading = ref(false)
    const fetchStatus = ref("success")

    async function fetchQuestions() {
        loading.value = true
        const { json } = await fetchJSONWithAuth('/api/v1/questions', user.get().value.accessToken)
        fetchStatus.value = json.status
        loading.value = false
        if (json.status === 'success') {
            questions.value = json.data.questions
        }
    }

    return { questions, fetchQuestions, fetchStatus, loading }
}