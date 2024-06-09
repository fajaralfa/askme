import { ref } from "vue"

const user = ref({
    id: null,
    email: null,
    photo: null,
    accessToken: null,
})

const key = 'askme_user'

function get() {
    if (user.value.accessToken == null) {
        const userJSON = localStorage.getItem(key)
        if (userJSON) {
            set(JSON.parse(userJSON))
            return user
        }

        return null
    }

    return user
}

function isLogin() {
    return user.value.accessToken != null
}

function set(userObj) {
    user.value.accessToken = userObj.accessToken
    user.value.id = userObj.user.id
    user.value.email = userObj.user.email
    user.value.photo = userObj.user.photo
    localStorage.setItem(key, JSON.stringify(userObj))
}

function forget() {
    user.value.accessToken = null
    user.value.id = null
    user.value.email = null
    user.value.photo = null
    localStorage.removeItem(key)
}

export { get, set, forget, isLogin }