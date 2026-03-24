# lock-hold
**Enhanced Commit Protocols for Low Latency Distributed Transactions**

* Author: Vijaya Krishna Namala
* Published In : International Journal of Innovative Research in Engineering & Multidisciplinary Physical Sciences (IJIRMPS)
* Publication Date: May 2022
* E-ISSN: 2349-7300
* Impact Factor : 9.907

**Abstract:**
Distributed transaction systems rely on commit protocols, but conventional approaches introduce high latency due to coordination overhead and multiple communication rounds. As systems scale, these delays reduce concurrency and increase sensitivity to network conditions. This work explores protocol level optimizations to restructure coordination and improve decision propagation while maintaining correctness. Experimental evaluation across cluster sizes demonstrates that optimized commit design can significantly reduce latency and improve scalability in distributed environments.

**Key Contributions**
* **Optimized Commit Coordination Design:**\
Developed an enhanced commit protocol that restructures coordination flow to reduce lock holding time and minimize blocking during distributed transaction execution.
* **Reduced Lock Retention Mechanism:**\
Introduced techniques for early lock release and parallel vote processing to limit resource holding duration and improve transaction concurrency.
* **Efficient Communication and Synchronization Model:**\
Designed a coordination approach that reduces unnecessary synchronization points and lowers communication overhead during commit decision propagation.
* **Experimental Evaluation Across Cluster Sizes and Contention Levels:**\
Conducted systematic analysis on clusters with 3, 5, 7, 9, and 11 nodes under varying contention conditions to evaluate lock hold time and scalability behavior.

**Relevance & Real World Impact**
* **Reduced Commit Latency :**\
The optimized protocol significantly lowers commit delay by minimizing coordination overhead and shortening lock holding duration in distributed transactions.
* **Improved Transaction Concurrency :**\
Faster lock release enables higher parallel transaction execution, reducing contention and improving overall system throughput.
* **Enhanced Scalability in Distributed Systems :**\
The protocol demonstrates controlled growth in coordination cost as cluster size increases, supporting efficient scaling of transactional systems.
* **Better Performance Under High Contention :**\
The approach maintains lower lock hold time even under heavy workload conditions, ensuring stable and responsive transaction processing.
* **Academic and Practical Contribution :**\
Provides a structured framework for improving commit protocol design, supporting further research and real world implementation in distributed databases and transaction platforms.

**Experimental Results (Summary)**:

  | Nodes | Conventional 2PC Protocol Lock Hold Time (ms) | Optimized Commit Transaction Protocol Lock Hold Time (ms) | Improvement (%) |
  |-------|-----------------------------------------------| ----------------------------------------------------------| ----------------|
  | 3     | 95                                            | 65                                                        | 31.58           |
  | 5     | 120                                           | 85                                                        | 29.17           |
  | 7     | 145                                           | 105                                                       | 27.59           |
  | 9     | 170                                           | 125                                                       | 26.47           |
  | 11    | 195                                           | 145                                                       | 25.64           |

**Citation** \
Enhanced Commit Protocols for Low Latency Distributed Transactions. \
Vijaya Krishna Namala \
International Journal of Innovative Research in Engineering & Multidisciplinary Physical Sciences \
E-ISSN- 2349-7300  \
**License** \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijirmps.org/ \
**Author Contact** \
**LinkedIn**: linkedin.com/in/vijaya-krishna-namala-a42b2958 | **Email**: vijaya.namala@gmail.com

