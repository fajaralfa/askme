import { ref } from "vue";
import * as user from '@/store/userStore'
import { useRouter } from "vue-router";

export default function useLogout() {
    const loading = ref(false)
    const router = useRouter()

    function submit() {
        user.forget()
        router.push({name: 'login'})
    }

    return { submit, loading }
}