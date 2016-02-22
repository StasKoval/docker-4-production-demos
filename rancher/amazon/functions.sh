# Creates a data volume:
function create_data_volume() {
  local JSON_FILE="$WORK_DIR/$1"
  echo "JSON File: $JSON_FILE"

  # aws ec2 create-volume \
  #   --volume-type io1 \
  #   --size 20 \
  #   --iops 600 \
  #   --availability-zone us-west-1a
}

# Creates a host on AWS using the given JSON file:
function create_host() {
  # Use the first argument as the json file path:
  local JSON_FILE=$1;

  # Use the second argument as the instance count, or default to 1:
  local INSTANCE_COUNT=${2:-1};

  # Create the instances using the AWS CLI:
  aws ec2 run-instances \
    --count $INSTANCE_COUNT \
    --cli-input-json $JSON_FILE \
    --output text \
    --query 'Instances[*].InstanceId'
}
