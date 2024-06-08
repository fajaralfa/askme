import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { fetchJSON } from '@/helper/fetch'
import useUserStore from './useUserStore'

export default function useLogin() {
    const router = useRouter()
    const user = useUserStore()

    const form = ref({
        email: '',
        password: '',
    })

    const loading = ref(false)
    const errMessage = ref('')

    async function submit(e) {
        e.preventDefault()
        loading.value = true
        const { json } = await fetchJSON('/api/v1/login', { method: 'post', body: form.value })
        loading.value = false
        if (json.status === 'success') {
            user.set(json.data)
            router.push({name: 'home'})
        } else {
            errMessage.value = json.message
        }
    }

    return {form, loading, errMessage, submit}
}
