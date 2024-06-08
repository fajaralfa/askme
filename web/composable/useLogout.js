import { ref } from "vue";
import useUserStore from "./useUserStore";
import { useRouter } from "vue-router";

export default function useLogout() {
    const loading = ref(false)
    const user = useUserStore()
    const router = useRouter()

    function submit() {
        user.forget()
        router.push({name: 'login'})
    }

    return { submit, loading }
}