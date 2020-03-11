<template>
  <meet-detail :is-edit="false" />
</template>

<script>
import MeetDetail from './components/MeetDetail'

export default {
  name: 'CreateMeet',
  components: { MeetDetail }
}
</script>
