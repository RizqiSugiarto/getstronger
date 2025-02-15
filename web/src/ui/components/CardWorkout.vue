<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useAuthStore } from '@/stores/auth.ts'
import { useTextareaAutosize } from '@vueuse/core'
import { useAlertStore } from '@/stores/alerts.ts'
import AppButton from '@/ui/components/AppButton.vue'
import { type DropdownItem } from '@/types/dropdown.ts'
import { FireIcon, UserCircleIcon } from '@heroicons/vue/24/solid'
import { formatToRelativeDateTime } from '@/utils/datetime.ts'
import DropdownButton from '@/ui/components/DropdownButton.vue'
import { deleteWorkout, postWorkoutComment } from '@/http/requests.ts'
import CardWorkoutComment from '@/ui/components/CardWorkoutComment.vue'
import CardWorkoutExercise from '@/ui/components/CardWorkoutExercise.vue'
import { type Workout, type WorkoutComment } from '@/proto/api/v1/workout_service_pb'
import AppCard from '@/ui/components/AppCard.vue'

const { input, textarea } = useTextareaAutosize()
const authStore = useAuthStore()
const alertStore = useAlertStore()
const workoutDeleted = ref(false)

const props = defineProps<{
  compact: boolean
  workout: Workout
}>()

const dropdownItems: Array<DropdownItem> = [
  { href: `/workouts/${props.workout.id}/edit`, title: 'Update Workout' },
  { func: () => onDeleteWorkout(), title: 'Delete Workout' },
]

const onDeleteWorkout = async () => {
  if (confirm('Are you sure you want to delete this workout?')) {
    await deleteWorkout(props.workout.id)
    alertStore.setErrorWithoutPageRefresh('Workout deleted')
    workoutDeleted.value = true
  }
}

const comments = ref<Array<WorkoutComment>>([])

onMounted(() => {
  comments.value = props.workout.comments
})

const postComment = async () => {
  const res = await postWorkoutComment(props.workout.id, input.value)
  if (!res) return
  if (!res.comment) return

  comments.value.push(res.comment)
  input.value = ''
}

const formatComment = computed(() => {
  if (comments.value.length === 1) return 'Comment'
  return `Comments`
})
</script>

<template>
  <AppCard v-if="!workoutDeleted" class="divide-y divide-gray-100">
    <div class="p-4">
      <div class="flex items-center justify-between">
        <UserCircleIcon class="size-10 text-gray-900" />
        <div class="w-full">
          <RouterLink :to="`/users/${props.workout.user?.id}`" class="font-semibold text-base mx-2">
            {{ props.workout.user?.firstName }} {{ props.workout.user?.lastName }}
          </RouterLink>
          <RouterLink :to="`/workouts/${workout.id}`" class="text-gray-500 text-sm">
            {{ formatToRelativeDateTime(props.workout.finishedAt) }}
          </RouterLink>
        </div>
        <DropdownButton v-if="workout.user?.id === authStore.userId" :items="dropdownItems" />
      </div>
    </div>
    <div class="p-4">
      <CardWorkoutExercise
        v-for="exerciseSet in workout.exerciseSets"
        :key="exerciseSet.exercise?.id"
        :exercise-id="exerciseSet.exercise?.id"
        :name="exerciseSet.exercise?.name"
        :sets="exerciseSet.sets"
      />
    </div>
    <div class="p-4 flex gap-x-1 font-medium text-sm uppercase justify-between">
      <div class="flex items-center">
        <RouterLink :to="`/workouts/${workout.id}`">
          {{ workout.comments.length }} {{ formatComment }}
        </RouterLink>
      </div>
      <div class="flex items-center">
        {{ workout.intensity.toLocaleString() }}
        <FireIcon class="size-4 text-orange-500 ml-1" />
      </div>
    </div>
    <div v-if="!compact" class="p-4">
      <CardWorkoutComment
        v-for="comment in comments"
        :key="comment.id"
        :user="comment.user"
        :timestamp="comment.createdAt"
        :comment="comment.comment"
      />
      <form @submit.prevent="postComment">
        <textarea
          ref="textarea"
          v-model="input"
          class="w-full border-2 border-gray-200 rounded-md focus:ring-indigo-500 focus:border-indigo-500 text-base min-h-20 py-3 mb-2 mt-2 resize-none overflow-hidden"
          placeholder="Write a comment..."
          required
        />
        <AppButton type="submit" colour="primary">Comment</AppButton>
      </form>
    </div>
  </AppCard>
</template>
