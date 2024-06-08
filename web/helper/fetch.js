export async function fetchJSON(url, options = { method : 'get', headers : {}, body : null }) {
    let res = null,
        json = null,
        err = null

    try {
        const init = {method: options.method, headers: options.headers}
        if (typeof options.body === 'object' && options.body != null) {
            init.body = JSON.stringify(options.body)
        }

        res = await fetch(url, init)
        json = await res.json()
    } catch (error) {
        err = error
    }

    return { res, json, err }
}

export async function fetchJSONWithAuth(url, token, options = { method: 'get', headers: {}, body: null }) {
    options.headers = {...options.headers, 'Authorization': `Bearer ${token}`}
    return await fetchJSON(url, options)
}