<template>
    <div class="container">
        <h3>包依赖关系图 / Package Dependency Graph</h3>
        <div id="mountNode"></div>
    </div>
</template>

<script>
import G6 from '@antv/g6';
import { mapGetters } from 'vuex';

export default {
    name: "LocalGraph",
    computed: {
        ...mapGetters(['graphData']),
    },
    watch: {
        graphData: {
            handler(newData) {
                this.destroyGraph();
                this.renderGraph(newData);
            },
            deep: false,
        },
    },
    methods: {
        renderGraph(graphData) {
            const container = document.getElementById('mountNode');
            const width = container.scrollWidth;
            const height = container.scrollHeight || 400;

            this.graph = new G6.Graph({
                container: 'mountNode', // 指定挂载容器
                width, // 图的宽度
                height, // 图的高度
                fitView: true,
                modes: {
                    default: ['zoom-canvas', 'drag-canvas', 'drag-node', 'activate-relations'],
                },
                layout: {
                    type: 'dagre',
                    rankdir: 'LR',
                    align: 'UL',
                    controlPoints: true,
                    nodesepFunc: () => 1,
                    ranksepFunc: () => 1,
                },
                animate: true,
                defaultNode: {
                    // 节点大小
                    type: 'rect',
                    size: [280, 30],
                    style: {
                        lineWidth: 2,
                        stroke: '#5B8FF9',
                        fill: '#C6E5FF',
                    },
                    // 节点上的标签文本配置
                    labelCfg: {
                        // 节点上的标签文本样式配置
                        style: {
                            fill: '#000', // 节点标签文字颜色
                            fontSize: 10,
                        },
                    },
                },
                defaultEdge: {
                    type: 'polyline',
                    size: 2,
                    color: '#000',
                    style: {
                        endArrow: {
                            path: 'M 0,0 L 8,4 L 8,-4 Z',
                            fill: '#e2e2e2',
                        },
                        radius: 20,
                    },
                },
            });
            this.graph.data(graphData); // 加载数据
            this.graph.render(); // 渲染
        },
        destroyGraph() {
            if (this.graph) {
                this.graph.destroy(); // 销毁图形实例
                this.graph = null;
            }
        }
    },
    mounted() {
        // 初始化图形
        this.renderGraph(this.graphData);
    }
}
</script>

<style scoped>
.container {
    margin-top: 20px;
}
</style>
