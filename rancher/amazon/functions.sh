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
  aws ec2 run-instances --cli-input-json $1
}
