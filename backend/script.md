try {
    const res = pm.response.json()
    pm.environment.set("token", res.data.token)
} catch {
    pm.environment.set("token", null)
}