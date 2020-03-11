<template>
  <meet-detail :is-edit="true" />
</template>

<script>
import MeetDetail from './components/MeetDetail'

export default {
  name: 'EditMeet',
  components: { MeetDetail }
}
</script>

