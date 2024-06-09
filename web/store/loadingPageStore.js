import { ref } from "vue"

export const loadingPageStore = ref(false)

export function isLoading() {
    return loadingPageStore
}

/**
 * 
 * @param {boolean} to 
 */
export function setLoading(to) {
    loadingPageStore.value = to
}