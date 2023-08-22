<template>
    <div class="container">
        <h3>包依赖关系图 / Package Dependency Graph</h3>
        <div id="mountNode"></div>
    </div>
</template>

<script>
import G6 from '@antv/g6';
import insertCss from 'insert-css';

// insertCss(`
//   .g6-component-toolbar li {
//     list-style-type: none !important;
//   }
// `);
insertCss(`
  .g6-component-tooltip {
    background-color: rgba(255, 255, 255, 0.8);
    padding: 0px 10px 24px 10px;
    box-shadow: rgb(174, 174, 174) 0px 0px 10px;
  }
`);

const data = {
    "nodes": [
        {
            "id": "github.com/pjimming/baize/common/errorx",
            "label": "github.com/pjimming/baize/common/errorx"
        },
        {
            "id": "github.com/pjimming/baize/common/httpresp",
            "label": "github.com/pjimming/baize/common/httpresp"
        },
        {
            "id": "github.com/pjimming/baize/core",
            "label": "github.com/pjimming/baize/core"
        },
        {
            "id": "github.com/pjimming/baize/core/helper",
            "label": "github.com/pjimming/baize/core/helper"
        },
        {
            "id": "github.com/pjimming/baize/core/internal/config",
            "label": "github.com/pjimming/baize/core/internal/config"
        },
        {
            "id": "github.com/pjimming/baize/core/internal/handler",
            "label": "github.com/pjimming/baize/core/internal/handler"
        },
        {
            "id": "github.com/pjimming/baize/core/internal/handler/basic",
            "label": "github.com/pjimming/baize/core/internal/handler/basic"
        },
        {
            "id": "github.com/pjimming/baize/core/internal/handler/local",
            "label": "github.com/pjimming/baize/core/internal/handler/local"
        },
        {
            "id": "github.com/pjimming/baize/core/internal/logic/basic",
            "label": "github.com/pjimming/baize/core/internal/logic/basic"
        },
        {
            "id": "github.com/pjimming/baize/core/internal/logic/local",
            "label": "github.com/pjimming/baize/core/internal/logic/local"
        },
        {
            "id": "github.com/pjimming/baize/core/internal/svc",
            "label": "github.com/pjimming/baize/core/internal/svc"
        },
        {
            "id": "github.com/pjimming/baize/core/internal/types",
            "label": "github.com/pjimming/baize/core/internal/types"
        }
    ],
    "edges": [
        {
            "source": "github.com/pjimming/baize/core",
            "target": "github.com/pjimming/baize/core/internal/handler"
        },
        {
            "source": "github.com/pjimming/baize/core",
            "target": "github.com/pjimming/baize/core/internal/svc"
        },
        {
            "source": "github.com/pjimming/baize/core/internal/handler/local",
            "target": "github.com/pjimming/baize/common/httpresp"
        },
        {
            "source": "github.com/pjimming/baize/core/internal/handler",
            "target": "github.com/pjimming/baize/core/internal/svc"
        },
        {
            "source": "github.com/pjimming/baize/core/internal/logic/local",
            "target": "github.com/pjimming/baize/common/errorx"
        },
        {
            "source": "github.com/pjimming/baize/common/httpresp",
            "target": "github.com/pjimming/baize/common/errorx"
        },
        {
            "source": "github.com/pjimming/baize/core",
            "target": "github.com/pjimming/baize/core/internal/config"
        },
        {
            "source": "github.com/pjimming/baize/core/helper",
            "target": "github.com/pjimming/baize/core/internal/types"
        },
        {
            "source": "github.com/pjimming/baize/core/internal/handler/basic",
            "target": "github.com/pjimming/baize/core/internal/svc"
        },
        {
            "source": "github.com/pjimming/baize/core/internal/handler",
            "target": "github.com/pjimming/baize/core/internal/handler/basic"
        },
        {
            "source": "github.com/pjimming/baize/core/internal/logic/basic",
            "target": "github.com/pjimming/baize/core/internal/svc"
        },
        {
            "source": "github.com/pjimming/baize/core/internal/logic/local",
            "target": "github.com/pjimming/baize/core/internal/svc"
        },
        {
            "source": "github.com/pjimming/baize/core/internal/svc",
            "target": "github.com/pjimming/baize/core/internal/config"
        },
        {
            "source": "github.com/pjimming/baize/core/internal/handler",
            "target": "github.com/pjimming/baize/core/internal/handler/local"
        },
        {
            "source": "github.com/pjimming/baize/core/internal/logic/local",
            "target": "github.com/pjimming/baize/core/helper"
        },
        {
            "source": "github.com/pjimming/baize/core/internal/logic/local",
            "target": "github.com/pjimming/baize/core/internal/types"
        },
        {
            "source": "github.com/pjimming/baize/core/internal/handler/local",
            "target": "github.com/pjimming/baize/common/errorx"
        },
        {
            "source": "github.com/pjimming/baize/core/internal/handler/local",
            "target": "github.com/pjimming/baize/core/internal/logic/local"
        },
        {
            "source": "github.com/pjimming/baize/core/internal/handler/local",
            "target": "github.com/pjimming/baize/core/internal/svc"
        },
        {
            "source": "github.com/pjimming/baize/core/internal/handler/local",
            "target": "github.com/pjimming/baize/core/internal/types"
        }
    ]
}


export default {
    name: "LocalGraph",

    mounted() {
        this.init();
    },
    methods: {
        init() {
            const container = document.getElementById('mountNode');
            const width = container.scrollWidth;
            const height = container.scrollHeight || 400;

            const tooltip = new G6.Tooltip({
                offsetX: 10,
                offsetY: 0,
                fixToNode: [1, 0],
                trigger: 'click',

                itemTypes: ['node'],
                getContent: (e) => {
                    const outDiv = document.createElement('div');
                    outDiv.style.width = 'fit-content';
                    outDiv.innerHTML = `
<h4>Custom Content</h4>
<ul>
    <li>Type: ${e.item.getType()}</li>
</ul>
<ul>
    <li>Label: ${e.item.getModel().label || e.item.getModel().id}</li>
</ul>
`;
                    return outDiv;
                },
            });

            const graph = new G6.Graph({
                container: 'mountNode', // 指定挂载容器
                width, // 图的宽度
                height, // 图的高度
                fitView: true,
                plugins: [tooltip],
                modes: {
                    default: ['zoom-canvas', 'drag-canvas', 'drag-node', 'activate-relations'],
                },
                // layout: {
                //     type: 'grid',
                //     begin: [20, 20],
                //     cols: 3,
                //     width: width - 20,
                //     height: height - 20,
                // },
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
            graph.data(data); // 加载数据
            graph.render(); // 渲染
        }
    },
    computed: {
        graphData() {
            return this.$store.state.graph;
        }
    },
}


</script>

<style scoped>
.container {
    margin-top: 20px;
}
</style>