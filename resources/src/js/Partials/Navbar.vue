<template>
    <div>
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="text-h6"> Application </v-list-item-title>
        </v-list-item-content>
      </v-list-item>

      <v-divider></v-divider>

      <v-list dense nav>
        <v-list-item v-for="item in items" :key="item.title" link>
          <v-list-item-icon>
            <img :src="item.icon" width="24" height="24" />
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>

        </v-list-item>
        <v-list-item>
          <v-list-item-icon>
            <v-icon>mdi-plus</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>Add New Wallet</v-list-item-title>
          </v-list-item-content>
          </v-list-item>
      </v-list>
    </div>
</template>

<script>
import binance from "@/img/organization/binance.png";
import eth from "@/img/organization/eth.png";
import ftx from "@/img/organization/ftx.png";

export default {
  data() {
    return {
      form: this.$inertia.form(),
      items: [
        { title: "Binance", icon: binance },
        { title: "FTX", icon: ftx },
        { title: "Eth Wallet", icon: eth },
      ],
    };
  },
};
</script>