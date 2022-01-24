<template>
  <div>
    <v-app-bar
        color="blue lighten-2"
        dense
        dark
    >
      <v-menu
          right
          bottom
      >
        <template v-slot:activator="{ on, attrs }">
          <v-btn
              icon
              v-bind="attrs"
              v-on="on"
          >
            <v-icon>mdi-dots-vertical</v-icon>
          </v-btn>
        </template>

        <v-list>
          <v-list-item
              v-for="(item,idx) in items"
              :key="idx"
              :to="item.to"
          >
            <v-list-item-title class="mr-15"> {{ item.title }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
      <v-toolbar-title>echo-vue-sample</v-toolbar-title>
    </v-app-bar>
  </div>
</template>

<script>
export default {
  name: "Header",
  data() {
    return {
      items: [
        {
          to: "/",
          title: "Home",
        },
        {
          to: "/test",
          title: "Test",
        },
      ]
    }
  }
}
</script>

<style scoped>

</style>
