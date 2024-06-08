import { ref } from "vue"
import user from '@/store/user';
import { fetchJSONWithAuth } from "@/helper/fetch";

export default function useGetQuestions() {
    const questions = ref([])
    const loading = ref(false)
    const fetchStatus = ref("success")
    
    async function fetchQuestions() {
        const {json} = await fetchJSONWithAuth('/api/v1/questions', user.value.accessToken)
        fetchStatus.value = json.status
        if (json.status === 'success') {
            questions.value = json.data.questions
        }
    }

    return { questions, fetchQuestions, fetchStatus, loading }
}