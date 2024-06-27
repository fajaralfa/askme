import { ref } from "vue"

const loadingPageStore = ref(false)

function isLoading() {
    return loadingPageStore
}

/**
 * 
 * @param {boolean} to 
 */
function setLoading(to) {
    loadingPageStore.value = to
}

export { isLoading, setLoading }