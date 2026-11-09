<script setup lang="ts">
import { ref, onMounted } from 'vue'

interface Info {
  name: string
  version: string
  language: string
  features: string[]
}

const info = ref<Info | null>(null)
const error = ref('')

onMounted(async () => {
  try {
    const res = await fetch('/api/info')
    if (!res.ok) throw new Error(`HTTP ${res.status}`)
    info.value = await res.json()
  } catch (e) {
    error.value = String(e)
  }
})
</script>

<template>
  <h1>You did it!</h1>
  <p v-if="error">Error: {{ error }}</p>
  <dl v-else-if="info">
    <dt>name</dt>
    <dd>{{ info.name }}</dd>
    <dt>version</dt>
    <dd>{{ info.version }}</dd>
    <dt>language</dt>
    <dd>{{ info.language }}</dd>
    <dt>features</dt>
    <dd>{{ info.features.join(', ') }}</dd>
  </dl>
  <p v-else>Loading…</p>
</template>
