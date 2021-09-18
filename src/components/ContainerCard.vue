<template>
  <el-card
    class="containers-card"
    shadow="hover"
  >
    <div
      slot="header"
      class="clearfix"
    >
      <h3 class="title">
        {{ container.Image }}
      </h3>
    </div>
    <p>{{ container.Names.join("|") }}</p>
    <div>{{ container.Status }}</div>
    <p class="time">
      {{ new Date(container.Created).toLocaleString() }}
    </p>
    <el-button
      v-if="container.State === 'running'"
      :class="container.State === 'running' ? 'btn-stop' : 'btn-restart'"
      class="bottom btn"
      style="float: right; padding: 3px 0"
      type="text"
      block
      :loading="loading"
      @click.native="stop(container.Id)"
    >
      停止
    </el-button>
    <el-button
      v-else
      :class="container.State === 'running' ? 'btn-stop' : 'btn-restart'"
      class="bottom btn"
      style="float: right; padding: 3px 0"
      type="text"
      block
      :loading="loading"
      @click.native="restart(container.Id)"
    >
      重启
    </el-button>
  </el-card>
</template>

<script>
import apis from "../router/request";

export default {
  name: "ContainerCard",
  props: {
    container: {
      type: Object,
      default: () => {},
    },
  },
  data() {
    return {
      loading: false,
    };
  },
  methods: {
    async stop(cId) {
      this.loading = true;
      try {
        await apis.stopContainer(cId);
        this.$parent.init();
      } finally {
        this.loading = false;
      }
    },
    async restart(cId) {
      this.loading = true;
      try {
        await apis.restartContainer(cId);
        this.$parent.init();
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>
<style>
.time {
  font-size: 13px;
  color: #999;
}

.bottom {
  margin-top: 13px;
  line-height: 12px;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}

.clearfix:after {
  clear: both;
}
</style>
<style scoped>
.containers-card {
  width: 450px;
  margin: 5px;
}
.title {
  text-align: center;
}
.btn {
  width: 100%;
  padding: 10px 0 !important;
  margin-bottom: 10px;
  color: #fff;
}

.btn:hover,
.btn:active,
.btn:focus {
  color: #fff;
}

.btn-stop {
  background: red;
}

.btn-stop:hover,
.btn-stop:active,
.btn-stop:focus {
  background: #ff4d00;
}

.btn-restart {
  background: lawngreen;
}

.btn-restart:hover,
.btn-restart:active.btn-restart:focus {
  background: #97fc00;
}
</style>
