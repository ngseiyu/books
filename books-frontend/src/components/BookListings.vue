<script setup lang="ts">
import { defineProps } from 'vue'
import type { Book } from '@/models/books'

const props = defineProps<{ book: Book }>()
</script>

<template>
  <section>
    <div>
      <RouterLink
        :to="'/books/' + props.book.id"
        class="block bg-white rounded-lg shadow-md hover:shadow-lg hover:bg-gray-50 transition-all duration-200 m-3 p-4 border border-gray-200"
      >
        <div class="m-3">
          <h3 class="text-lg font-semibold text-gray-800 mb-1">{{ props.book.title }}</h3>
          <h3 class="text-sm text-gray-600 mb-1">{{ props.book.author }}</h3>
          <h3 class="text-sm text-gray-500">{{ props.book.genre }}</h3>
        </div>
      </RouterLink>
    </div>
  </section>
</template>
