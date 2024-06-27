import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { fetchJSON } from '@/helper/fetch'

export default function useRegister() {
    const router = useRouter()

    const form = ref({
        email: '',
        password: '',
    })

    const loading = ref(false)
    const errMessage = ref('')

    async function submit(e) {
        e.preventDefault()
        loading.value = true
        const { json } = await fetchJSON('/api/v1/register', { method: 'post', body: form.value })
        loading.value = false
        if (json.status === 'success') {
            router.push({ name: 'login' })
        } else {
            errMessage.value = json.message
        }
    }

    return { form, loading, errMessage, submit }
}
