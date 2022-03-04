<template>
  <v-app id="inspire">
    <!-- <v-system-bar app>
      <v-spacer></v-spacer>

      <v-icon>mdi-square</v-icon>

      <v-icon>mdi-circle</v-icon>

      <v-icon>mdi-triangle</v-icon>
    </v-system-bar> -->

    <!-- <v-app-bar app> -->
      <!-- <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon> -->

      <!-- <v-toolbar-title>Application</v-toolbar-title>
    </v-app-bar> -->

    <v-navigation-drawer app permanent>
      <!--  -->
      <Navbar />
    </v-navigation-drawer>

    <v-main class="grey lighten-2">
      <v-container>
        <Flash
          v-model="$page.props.flash.message"
          v-if="$page.props.flash && $page.props.flash.message"
        />
                <slot />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import Navbar from "@/js/Partials/Navbar";
import Flash from "@/js/Shared/Flash";

export default {
  components: {
    Navbar,
    Flash,
  },
  data: () => ({ drawer: null }),
};
</script>