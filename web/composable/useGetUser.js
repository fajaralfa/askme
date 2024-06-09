import { ref } from "vue"
import { fetchJSON } from '@/helper/fetch';

export default function useGetUser() {
    const user = ref({})
    const loading = ref(false)

    async function getUser() {
        loading.value = true
        const { json } = await fetchJSON(`/api/v1/users/${props.email}`)
        loading.value = false
        if (json.data) {
            user.value = json.data.user
        }
    }

    return { user, getUser, loading }
}