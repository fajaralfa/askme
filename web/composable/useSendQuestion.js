import { fetchJSON } from '@/helper/fetch'
import { ref } from 'vue'
import { useRoute } from 'vue-router'

export default function useSendQuestion() {
    const route = useRoute()

    const form = ref({
        question: '',
        targetEmail: route.params.email,
    })
    const loading = ref(false)
    const status = ref('')

    async function submit(e) {
        e.preventDefault()
        loading.value = true
        const { json } = await fetchJSON('/api/v1/questions', {
            method: 'post',
            body: form.value,
        })
        status.value = json.status
        loading.value = false
    }

    return { form, loading, submit, status }
}
