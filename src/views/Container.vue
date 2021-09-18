<template>
  <div id="docker-containers">
    <ContainerCard
      v-for="container in containers"
      :key="container.Id"
      class="containers-div"
      :container="container"
    />
    <el-drawer
      title="筛选"
      :with-header="false"
      :visible.sync="drawer"
      :direction="'btt'"
      size="100px"
    >
      <div class="filter-box">
        <el-checkbox v-model="option.all">
          显示全部
        </el-checkbox>
        <el-checkbox v-model="option.size">
          显示大小
        </el-checkbox>
        <el-input-number
          v-model="option.limit"
          :min="0"
          :max="100"
          size="mini"
          label="显示数量"
        />
      </div>
    </el-drawer>
    <el-button
      class="filter-btn"
      type="primary"
      @click.native="drawer = true"
    >
      筛选
    </el-button>
  </div>
</template>

<script>
import apis from "../router/request";
import ContainerCard from "../components/ContainerCard";

export default {
  name: "Container",
  mounted() {
    this.init();
  },
  components: { ContainerCard },
  data() {
    return {
      drawer: false,
      containers: [],
      option: {
        all: true,
        size: true,
        limit: 0,
      },
    };
  },
  computed: {
    all() {
      return this.option.all;
    },
    size() {
      return this.option.size;
    },
    limit() {
      return this.option.limit;
    },
  },
  watch: {
    all() {
      this.init();
    },
    size() {
      this.init();
    },
    limit() {
      this.init();
    },
  },
  methods: {
    init() {
      this.$nextTick(async () => {
        let res = await apis.ListContainers(this.option);
        this.containers = res.data.result.containers;
      });
    },
  },
};
</script>

<style scoped>
#docker-containers {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}
.filter-box {
  margin: 10px;
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
}
.filter-box label,
.filter-box > input {
  margin: 10px;
}
.filter-btn{
  position: fixed;
  bottom: -10px;
}
</style>
