# Golang FaceID API
This API aims to create a backend infrastructure to allow user authentication by Face Recognition. The API receives a FRV (Face Representation Vector) from remote client and verifies if the given object matchs an existing FRV in the vector database. The API simply grants or deny access to the client.

This API can be used from any client that can send a HTTP request. The API is written in Golang and uses the REST architecture. The client can be written in any language, and can come from any platform (web, mobile, desktop, etc). This facilitates the integration of Face Recognition in any application. The process of getting the FRV from the client is not part of this API. 

## Face Recognition Technology
Face recognition is a technology that involves identifying and verifying individuals by analyzing and comparing their facial features. It's a subset of biometric identification, and it has various applications, including security systems, authentication, and user access control. Here's a simplified overview of how face recognition works:

**Face Detection:** The first step is to locate and extract faces from an image or video frame. This is done using face detection algorithms, which identify potential face regions in the input data.

**Face Alignment:** Once a face is detected, the algorithm might perform face alignment, ensuring that the face is in a standardized position. This helps in normalizing the features for accurate comparison.

**Feature Extraction:** This is a crucial step where distinctive features of the face are extracted and turned into a numerical representation. These features might include the distances between key facial landmarks, the angles between certain points, and other data that can differentiate one face from another. Popular techniques for feature extraction include Local Binary Patterns (LBP), Histogram of Oriented Gradients (HOG), and more recently, deep learning-based methods.

**Feature Encoding:** The extracted features are then converted into a compact representation that can be easily compared with other face representations. This often involves reducing the dimensionality of the features while preserving their discriminatory information.

**Face Database:** In your scenario, you would have a database containing the face representations (or embeddings) of authorized users. This database could be created by collecting and processing a set of images of each user during enrollment.

**Face Matching:** When a user attempts to gain access, their submitted face is processed through the same pipeline: detection, alignment, feature extraction, and encoding. The resulting face representation is then compared with the representations in your face database using a distance metric like Euclidean distance or cosine similarity. The idea is to find the closest match.

**Thresholding:** The comparison produces a similarity score. You would set a threshold value above which the submitted face is considered a match with a face in the database. The threshold helps balance false positives and false negatives.

**Decision Making:** Based on the similarity score and the threshold, a decision is made whether to grant access or not. If the score is above the threshold and within an acceptable range, access is granted. Otherwise, access is denied.

**Continuous Learning:** To improve recognition accuracy over time, some systems implement continuous learning. This involves periodically updating the face representations in the database using new images of users. This can help account for changes in appearance due to factors like aging or facial hair.


