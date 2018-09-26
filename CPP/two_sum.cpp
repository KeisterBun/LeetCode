/* Written by Keith Drew
 * This is a solution to LeetCode problem #1: "Two Sum," written in C++, 2018
 * 
 * The answer could be better. Some improvements:
 * - Handle the case of no solution better.
 * - Handle the case of multiple solutions.
 * - Optimize the code.
 * - Read input for the vector from STDIN to more quickly test many inputs. 
 * - ??
 */
#include <iostream>
#include <vector>

std::vector<int> twoSum(std::vector<int> nums, int target);

int main() {
  std::vector<int> nums;
  int target = 26;

  // Add the data to our vector.
  nums.push_back(2);
  nums.push_back(7);
  nums.push_back(11);
  nums.push_back(15);

  // Call twoSum
  std::vector<int> indices = twoSum(nums, target);

  // Check if the number of outputs is correct. Print results.
  if (indices.size() != 2) {
    std::cout << "Failure. twoSum returned the wrong number of elements." << std::endl;
  } else {
    std::cout << "Success - indices are: " << indices.at(0) << ", " << indices.at(1) << std::endl;
  }
  
  return 0;  
}

/* The main code needed by LeetCode, pretty close to their provided prototype.
 *
 */
std::vector<int> twoSum(std::vector<int> nums, int target) {
  std::vector<int> indices;

  for (int i = 0; i < nums.size(); i++) {
    for (int j = i+1; j < nums.size(); j++) {
      if (nums.at(i) + nums.at(j) == target) {
	indices.push_back(i);
	indices.push_back(j);
	return indices;
      }
    }
  }
  return indices;
}
